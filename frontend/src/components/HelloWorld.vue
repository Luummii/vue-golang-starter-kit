<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <button @click="setData">ОТПРАВИТЬ</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      First: 'Andrey',
      Last: 'Zimin'
    }
  },
  methods: {
    setData () {
      if (this.First != '' && this.Last != '') {
        axios.post('/users/add', { First: this.First, Last: this.Last, Password: '12345', Email: 'email@reg.qwe' }).then((res) => { 
          console.log('res.data = ', res.data)              
        }).catch((err) => {
          console.log(err)
        })
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
</style>
