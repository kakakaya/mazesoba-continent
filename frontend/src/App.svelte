<script>
  import logo from './assets/images/logo-universal.png'
  import {Chikuwa, Greet, Post} from '../wailsjs/go/main/App.js'

  let resultText = "Please enter your name below 👇"
  let text

  function greet() {
    Greet(text).then(result => resultText = result)
  }

 function post() {
     Post(text).then(result => resultText = result)
 }

 function chikuwa() {
     Chikuwa().then(result => resultText = result)
 }

</script>

<main>
    <!-- <img alt="Wails logo" id="logo" src="{logo}"> -->
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <textarea autocomplete="off" bind:value={text} id="name" />
    <button class="btn" on:click={greet}>Greet</button>
    <button class="btn" on:click={post}>Post</button>
    <button class="btn" on:click={chikuwa}>ちくわ。</button>
  </div>
</main>

<style>

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  textarea {
    color: white;
    background: transparent;
    resize: none;
    border: 0 none;
    width: 100%;
  }

</style>
