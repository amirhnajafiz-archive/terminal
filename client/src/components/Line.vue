<script setup>
defineProps({
  data: {
    type: String,
    required: true
  },
  type: {
    type: String,
  }
})
</script>

<template>
  <div v-bind:class="this.getType()">
    {{ data }}
  </div>
</template>

<script>
export default {
  name: "Line",
  methods: {
    getType() {
      switch (this.type) {
        case "danger":
          return "danger"
        case "success":
          return "success"
        default:
          return "normal"
      }
    }
  }
}
</script>

<style scoped>
.normal {
  color: #222222;
}

.danger {
  color: #aa1111;
}

.success {
  color: #008800;
}
</style>