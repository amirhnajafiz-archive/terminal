<template>
  <div style="width: 100%">
    <input
        id="cmd"
        class="normal"
        type="text"
        v-model="this.in"
        placeholder="..."
        @keydown.enter="click"
    />
  </div>
</template>

<script>
export default {
  name: "Input",
  data() {
    return {
      in: ""
    }
  },
  methods: {
    click() {
      if (this.in === "clear") {
        this.$emit('clear')
      } else {
        this.$emit('submit', this.in)
      }
      this.in = ""
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