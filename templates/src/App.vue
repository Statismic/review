<template>
  <ion-app>
    <ion-content padding>
      <div id="container" v-if="!uuid">
        <ion-item lines="none">
          <ion-label>I want to review</ion-label>
          <ion-input
            inputmode="number"
            placeholder="PR number"
            class="input"
            @ionChange="onChangeID($event)"
          ></ion-input>
        </ion-item>
        <ion-item lines="none">
          <ion-label>of</ion-label>
          <ion-input
            inputmode="url"
            type="url"
            placeholder="Repo's url"
            class="input"
            @ionChange="onChangeURL($event)"
          ></ion-input>
        </ion-item>

        <ion-button v-show="validated" color="primary" @click="review()">Review</ion-button>
      </div>

      <div id="container" v-else>
        <viewer :url="`/${uuid}/master/`" title="Master" style="width: 50%; height: 100%;"></viewer>
        <viewer :url="`/${uuid}/pr/`" :title="`PR #${prID}`" style="width: 50%; height: 100%;"></viewer>
      </div>

      <ion-spinner name="crescent" color="primary" id="loading" v-show="loading"></ion-spinner>
      <div id="overlay" v-show="loading"></div>
    </ion-content>

    <ion-fab vertical="bottom" horizontal="end" slot="fixed">
      <ion-fab-button @click="reset()">
        <ion-icon name="refresh"></ion-icon>
      </ion-fab-button>
    </ion-fab>
  </ion-app>
</template>

<script>
import Viewer from "@/components/Viewer";

export default {
  name: "app",
  components: {
    Viewer
  },
  data() {
    return {
      url: "",
      prID: -1,
      loading: false,
      uuid: ""
    };
  },
  computed: {
    validated() {
      const regex = /^https?:\/\/.+\.git$/;
      return regex.test(this.url) && this.prID > 0;
    }
  },
  methods: {
    onChangeID(e) {
      this.prID = parseInt(e.target.value);
    },
    onChangeURL(e) {
      this.url = e.target.value;
    },
    async review() {
      try {
        this.loading = true;
        const resp = await fetch("/clone", {
          method: "POST",
          mode: "cors",
          headers: {
            "Content-Type": "text/plain"
          },
          body: JSON.stringify({
            url: this.url,
            id: this.prID
          })
        });

        this.uuid = await resp.text();
      } finally {
        this.loading = false;
      }
    },
    reset() {
      this.url = "";
      this.prID = -1;
      this.uuid = "";
    }
  }
};
</script>

<style>
#container {
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: row;
}

.input {
  display: inline;
  text-align: center;
  white-space: nowrap;
  border: 0;
  outline: 0;
  background: transparent;
  border-bottom: 1px solid black;
}

#loading {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 64px;
  height: 64px;
}

#overlay {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  z-index: 9999;
  background-color: rgba(0, 0, 0, 0.2);
}
</style>
