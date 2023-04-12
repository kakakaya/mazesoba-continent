package main

import (
	"context"
	"fmt"
	"log"
	"time"

	comatproto "github.com/bluesky-social/indigo/api/atproto"
	appbsky "github.com/bluesky-social/indigo/api/bsky"
	lexutil "github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/gen2brain/beeep"
)

func PostFeed(xrpcc *xrpc.Client, text string) (string, error) {
	// xrpcc, err := cliutil.GetXrpcClient(ctx, true)
	err := beeep.Notify("まぜそば大陸", fmt.Sprintf("PostFeed: posting %s", text), "")
	if err != nil {
		log.Fatalln(err)
	}

	return "<MOCK URI>", nil // TEMP

	resp, err := comatproto.RepoCreateRecord(context.TODO(), xrpcc, &comatproto.RepoCreateRecord_Input{
		Collection: "app.bsky.feed.post",
		Repo:       xrpcc.Auth.Did,
		Record: &lexutil.LexiconTypeDecoder{&appbsky.FeedPost{
			Text:      text,
			CreatedAt: time.Now().Format("2006-01-02T15:04:05.000Z"),
		}},
	})
	if err != nil {
		return "", fmt.Errorf("failed to create post: %w", err)
	}
	log.Println(resp.Cid, resp.Uri)
	return resp.Uri, nil
}
