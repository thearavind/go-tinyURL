<template>
  <v-container fluid="fluid" class="text-xs-center">
    <v-container id="home" fluid="fluid" class="text-xs-center">
      <v-layout row wrap>
        <v-flex xs10 offset-xs1>
          <v-card class="base-card black--text">
            <v-layout row wrap>
              <v-flex xs10 offset-xs1>
                <div class="headline">
                  <p class="tiny-header">
                    A simple but powerful app to shorten URL's, written in Golang
                    <a href="https://github.com/thearavind/go-tinyURL">
                      <img src="../assets/fork-me.png"
                           style="top: 8px; position:relative; cursor: pointer;" alt="star"></a>
                  </p>
                </div>
                <div>
                  <p>Built with <a href="https://www.postgresql.org/">PostgreSQL</a>,<a
                    href="https://gin-gonic.github.io/gin/">Gin Framework</a>, <a href="https://golang.org/">Golang</a>, <a
                    href="https://vuejs.org/">VueJS</a></p>
                </div>
                <v-text-field class="url-input" color="url-input grey darken-4 white--text"
                              v-model="inputURL"
                              label="Enter the URL to be shorten here"
                ></v-text-field>
              </v-flex>
            </v-layout>
            <v-layout row wrap>
              <v-flex xs10 offset-xs1>
                <v-btn color="grey darken-4 white--text" @click="shortenUrl()">Shorten</v-btn>
              </v-flex>
            </v-layout>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
    <v-snackbar
      :timeout="6000"
      :top="true"
      :multi-line="false"
      :vertical="false"
      v-model="snackbar">
      <span id="url">{{tinyURL}}</span>
      <v-btn flat icon color="green lighten-3" @click.native="snackbar = false">
        <v-icon>clear</v-icon>
      </v-btn>
      <v-btn v-if="showCopy" flat icon color="green lighten-3" class="btn butts" @click="copyToClipboard()"
             data-clipboard-target="#url">
        <v-icon>content_copy</v-icon>
      </v-btn>
    </v-snackbar>
  </v-container>
</template>

<script>
  import axios from 'axios'
  import Clipboard from 'clipboard'

  export default {
    data () {
      return {
        inputURL: '',
        baseURL: 'http://go-tiny.herokuapp.com/',
        snackbar: false,
        showCopy: false,
        tinyURL: ''
      }
    },
    methods: {
      shortenUrl () {
        console.log('inputURL', this.isUrlValid(this.inputURL))
        if (this.isUrlValid(this.inputURL)) {
          axios.post(`/short`, {
            'address': this.inputURL
          })
            .then(response => {
              this.snackbar = true
              this.tinyURL = this.baseURL + response.data.tiny_url.hash
              this.showCopy = true
            })
            .catch(e => {
              this.snackbar = true
              this.tinyURL = 'Error! Please try again'
              this.showCopy = false
            })
        } else {
          this.snackbar = true
          this.tinyURL = 'Please enter a valid URL'
          this.showCopy = false
        }
      },
      copyToClipboard () {
        let clipboard = new Clipboard('.butts')
        clipboard.on('success', e => {
          e.clearSelection()
        })
      },

      isUrlValid (userInput) {
        let res = userInput.match(/^(http(s)?:\/\/.)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&\/=]*)$/g)
        return res !== null
      }
    }
  }
</script>

<style scoped>
  h1, h2 {
    font-weight: normal;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    display: inline-block;
    margin: 0 10px;
  }

  a {
    color: #42b983;
  }

  .base-card {
    opacity: 0.8;
    padding: 50px 30px;
    background-color: #FAFAFA !important;
  }

  .headline {
    text-align: center;
    margin-bottom: 20px;
  }

  .tiny-header {
    margin-bottom: 35px;
    font-size: 28px;
  }

  .url-input {
    margin-top: 35px;
  }

  #home {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
  }
</style>
