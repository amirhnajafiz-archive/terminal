<template>
  <div class="container">
    <Line
        v-for="line in this.lines"
        :data="'$/ ' + line"
    >
    </Line>
  </div>
</template>

<script>
import Line from "@/components/Line.vue";

export default {
  name: "Terminal",
  components: {Line},
  data() {
    return {
      lines: [
          "hello",
      ]
    }
  },
  methods: {
    addCommand(str) {
      this.lines.push(str)
    }
  }
}
</script>

<style scoped>
.container {
  width: 100%;
  padding: 5px 10px;
  background: #ffffff;
  color: #222222;
  border-radius: 4px;

  max-height: 500px;
  overflow: scroll;
}
</style>