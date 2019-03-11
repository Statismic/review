<template>
  <ion-card id="container">
    <ion-chip id="title">
      <ion-label color="primary">{{ title }}</ion-label>
    </ion-chip>

    <!-- fab placed to the bottom and start and on the bottom edge of the content overlapping footer with a list to the right -->
    <ion-fab vertical="top" horizontal="start" slot="fixed">
      <ion-fab-button>
        <ion-icon name="settings"></ion-icon>
      </ion-fab-button>
      <ion-fab-list side="end">
        <ion-fab-button @click="back()">
          <ion-icon name="arrow-round-back"></ion-icon>
        </ion-fab-button>
        <ion-fab-button @click="forward()">
          <ion-icon name="arrow-round-forward"></ion-icon>
        </ion-fab-button>
        <ion-fab-button>
          <ion-icon name="refresh"></ion-icon>
        </ion-fab-button>
      </ion-fab-list>
    </ion-fab>
    <iframe :src="url" class="viewer" ref="viewer"></iframe>
  </ion-card>
</template>

<script>
export default {
  props: {
    url: String,
    title: String
  },
  methods: {
    back() {
      this.$refs.viewer.contentWindow.history.go(-1);
    },
    forward() {
      this.$refs.viewer.contentWindow.history.go(1);
    }
  }
};
</script>

<style scoped>
#container {
  position: relative;
}

#title {
  position: absolute;
  left: 50%;
  top: 20px;
  transform: translateX(-50%);
}

.viewer {
  width: 100%;
  height: 100%;
  padding-top: 64px;
}
</style>
