package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	comatproto "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	appbsky "github.com/bluesky-social/indigo/api/bsky"
	lexutil "github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
)

func addLink(xrpcc *xrpc.Client, post *bsky.FeedPost, link string) {
	doc, err := goquery.NewDocument(link)
	var title string
	var description string
	var imgURL string

	if err == nil {
		title = doc.Find(`title`).Text()
		description, _ = doc.Find(`meta[property="description"]`).Attr("content")
		imgURL, _ = doc.Find(`meta[property="og:image"]`).Attr("content")
		if title == "" {
			title, _ = doc.Find(`meta[property="og:title"]`).Attr("content")
			if title == "" {
				title = link
			}
		}
		if description == "" {
			description, _ = doc.Find(`meta[property="og:description"]`).Attr("content")
			if description == "" {
				description = link
			}
		}
		post.Embed.EmbedExternal = &bsky.EmbedExternal{
			External: &bsky.EmbedExternal_External{
				Description: description,
				Title:       title,
				Uri:         link,
			},
		}
	} else {
		post.Embed.EmbedExternal = &bsky.EmbedExternal{
			External: &bsky.EmbedExternal_External{
				Uri: link,
			},
		}
	}
	if imgURL != "" && post.Embed.EmbedExternal != nil {
		resp, err := http.Get(imgURL)
		if err == nil && resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			b, err := io.ReadAll(resp.Body)
			if err == nil {
				resp, err := comatproto.RepoUploadBlob(context.TODO(), xrpcc, bytes.NewReader(b))
				if err == nil {
					post.Embed.EmbedExternal.External.Thumb = &lexutil.LexBlob{
						Ref:      resp.Blob.Ref,
						MimeType: http.DetectContentType(b),
						Size:     resp.Blob.Size,
					}
				}
			}
		}
	}
}

func BskyFeedPost(xrpcc *xrpc.Client, text string) (string, error) {
	// Post given text to Bluesky, with app.bsky.feed.post.
	if text == "" || strings.TrimSpace(text) == "" {
		return "", fmt.Errorf("no text specified")
	}

	post := &appbsky.FeedPost{
		Text:      text,
		CreatedAt: time.Time.UTC(time.Now()).Format("2006-01-02T15:04:05.000Z"),
	}

	for _, entry := range extractLinksBytes(text) {
		post.Facets = append(post.Facets, &appbsky.RichtextFacet{
			Features: []*appbsky.RichtextFacet_Features_Elem{
				{
					RichtextFacet_Link: &appbsky.RichtextFacet_Link{
						Uri: entry.text,
					},
				},
			},
			Index: &appbsky.RichtextFacet_ByteSlice{
				ByteStart: entry.start,
				ByteEnd:   entry.end,
			},
		})
		if post.Embed == nil {
			post.Embed = &appbsky.FeedPost_Embed{}
		}
		if post.Embed.EmbedExternal == nil {
			addLink(xrpcc, post, entry.text)
		}
	}

	for _, entry := range extractMentionsBytes(text) {
		profile, err := appbsky.ActorGetProfile(context.TODO(), xrpcc, entry.text)
		if err != nil {
			return "", err
		}
		post.Facets = append(post.Facets, &appbsky.RichtextFacet{
			Features: []*appbsky.RichtextFacet_Features_Elem{
				{
					RichtextFacet_Mention: &appbsky.RichtextFacet_Mention{
						Did: profile.Did,
					},
				},
			},
			Index: &appbsky.RichtextFacet_ByteSlice{
				ByteStart: entry.start,
				ByteEnd:   entry.end,
			},
		})
	}

	for _, entry := range extractTagsBytes(text) {
		post.Facets = append(post.Facets, &appbsky.RichtextFacet{
			Features: []*appbsky.RichtextFacet_Features_Elem{
				{
					RichtextFacet_Tag: &appbsky.RichtextFacet_Tag{
						Tag: entry.text,
					},
				},
			},
			Index: &appbsky.RichtextFacet_ByteSlice{
				ByteStart: entry.start,
				ByteEnd:   entry.end,
			},
		})
	}

	resp, err := comatproto.RepoCreateRecord(context.TODO(), xrpcc, &comatproto.RepoCreateRecord_Input{
		Collection: "app.bsky.feed.post",
		Repo:       xrpcc.Auth.Did,
		Record: &lexutil.LexiconTypeDecoder{
			Val: post,
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to create post: %w", err)
	}
	return resp.Uri, nil
}
