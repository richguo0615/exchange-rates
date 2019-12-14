<template>
  <div id="app">
    <img src="./assets/logo.png">
    <router-view/>
    <div class="bg-left"></div>
    <div class="bg-right"></div>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
#app > img {
  margin-top: 50px;
  width: 400px;
  max-width: 60%;
}
.bg-left, .bg-right {
  position: fixed;
  z-index: -1;
  width: 20vw;
  height: 100vh;
  background-size: contain;
}
.bg-left {
  left: 0;
  top: 0;
  background-image: url("../static/images/bg-left.png");
}
.bg-right {
  right: 0;
  top: 0;
  background-image: url("../static/images/bg-right.png");
}
</style>
