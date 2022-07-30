<template>
  <div class="container">
    <div class="flex" v-for="line in this.lines">
      <span class="right">
        $/
      </span>
      <Line
          :data="line.msg"
          :type="line.typ"
      >
      </Line>
    </div>
    <div class="flex">
      <span class="right">
        $/
      </span>
      <Input
          @submit="addCommand"
          @clear="clear"
      />
    </div>
  </div>
</template>

<script>
import Line from "@/components/Line.vue";
import Input from "@/components/Input.vue";

export default {
  name: "Terminal",
  components: {Input, Line},
  data() {
    return {
      lines: [],
    }
  },
  methods: {
    addCommand(str, type='normal') {
      this.lines.push({
        msg: str,
        typ: type
      })
    },
    clear() {
      this.lines = []
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

.flex {
  display: flex;
  flex-direction: row;
}

.right {
  margin-right: 10px;
}
</style>