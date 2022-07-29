<script setup>
import Terminal from "@/components/Terminal.vue";
</script>

<template>
  <main>
    <Terminal />
  </main>
</template>

<script>
export default {
  name: "App",
  mounted() {
    document.getElementById("cmd").focus()
    document.body.addEventListener("click", function () {
      document.getElementById("cmd").focus()
    })
  },
}
</script>

<style scoped>
main {
  width: 60%;
  margin: 20px auto;
}
</style>
