<template>
        <v-shell
            :banner="banner"
            :shell_input="command"
            :commands="commands"
        ></v-shell>
</template>

<script>
export default {
    name: "Terminal",
    data() {
        return {
            command: "",
            banner: {
                header: "Terminal Project",
                subHeader: "Use terminal project to communicate with our server.",
                emoji: {
                    first: "*",
                    second: "*",
                    time: Date.now()
                },
                img: null,
                sign: "user$",
            },
            commands: [
                {
                    name: "help",
                    desc: "Enter help to see the manual",
                    get() {
                        return "Help manual"
                    },
                }
            ]
        }
    }
}
</script>

<style scoped>

</style>
