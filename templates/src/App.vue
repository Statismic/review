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

      <div v-else style="display: flex; flex-direction: column; height: 100%; align-items: center;">
        <ion-select value="both" @ionChange="onVisibilityChange($event)">
          <ion-select-option value="master">Master</ion-select-option>
          <ion-select-option value="pr">PR</ion-select-option>
          <ion-select-option value="both">Both</ion-select-option>
        </ion-select>
        <div id="container">
          <viewer v-show="isMasterVisible" :url="`/${uuid}/master/`" title="Master"></viewer>
          <viewer v-show="isPRVisible" :url="`/${uuid}/pr/`" :title="`PR #${prID}`"></viewer>
        </div>
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
      uuid: "",
      isMasterVisible: true,
      isPRVisible: true
    };
  },
  computed: {
    validated() {
      const regex = /^https?:\/\/.+\.git$/;
      return regex.test(this.url) && this.prID > 0;
    }
  },
  methods: {
    onVisibilityChange(e) {
      const visibility = e.target.value;
      switch (visibility) {
        case "master":
          this.isMasterVisible = true;
          this.isPRVisible = false;
          break;
        case "pr":
          this.isMasterVisible = false;
          this.isPRVisible = true;
          break;
        default:
          this.isMasterVisible = true;
          this.isPRVisible = true;
      }
    },
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
      this.isMasterVisible = true;
      this.isPRVisible = true;
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
