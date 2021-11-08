<template>
    <div>
        <h1>
            Hi world!
        </h1>
        <vue-terminal title="Online Terminal" />
    </div>
</template>

<script>
    import VueTerminal from 'vue-terminal';
    export default {
        name: "HelloWorld",
        components: {
            VueTerminal
        }
    }
</script>

<style scoped>

</style>
