<script setup>
defineProps({
  data: {
    type: String,
    required: true
  }
})
</script>

<template>
  <div>
    {{ data }}
  </div>
</template>

<script>
export default {
  name: "Line"
}
</script>

<style scoped>

</style>