<template>
  <div id="app"> 
    <router-view/>
  </div>
</template>

<script>
const inBrowser = typeof window !== 'undefined'
if (inBrowser) {
  var semantic = require('../semantic/dist/semantic.js')
}

export default {
  name: 'App',
  data () {
    return {
    }
  }
}
</script>

<style lang="scss">

</style>
