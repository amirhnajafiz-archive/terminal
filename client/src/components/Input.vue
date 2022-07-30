<template>
  <div style="width: 100%">
    <input
        id="cmd"
        class="normal"
        type="text"
        v-model="this.in"
        placeholder="..."
        @keydown.enter="click"
        @keydown.up="up"
        @keydown.down="down"
    />
  </div>
</template>

<script>
export default {
  name: "Input",
  data() {
    return {
      in: "",
      messageHistory: [],
      index: 0,
    }
  },
  methods: {
    click() {
      if (this.in === "clear") {
        this.$emit('clear')
      } else {
        this.messageHistory.push(this.in)
        this.fetch(this.in)
        this.$emit('submit', this.in)
      }
      this.in = ""
      this.index = 0
    },
    up() {
      if (this.index !== this.messageHistory.length) {
        this.index++
      }
      this.in = this.messageHistory[this.messageHistory.length - this.index]
    },
    down() {
      if (this.index !== 0) {
        this.index--
      }
      this.in = this.messageHistory[this.messageHistory.length - this.index]
    },
    async postData(url = '', data = {}) {
      const response = await fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
          'Content-Type': 'application/json'
        },
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(data)
      });
      return response.json();
    },
    fetch(command) {
      this.postData('http://localhost:3000/api/cmd', {command: command})
          .then((data) => {
            console.log(data)
          })
    }
  }
}
</script>

<style scoped>
.normal {
  width: 100%;
  border: 0 solid black;
  outline: none;
}
</style>