// TODO: This is most likely broken now. Need to fix it once DB is plugged.
$(() => {
    document.querySelector(".copy_blueprintstr").addEventListener("click", event => {
        console.log(event.target.toString() + " .bpstr")
        var copyText = document.querySelector(event.target.toString() + " .bpstr")
        copyText.select()
        console.log("Copied", copyText.value)
        document.execCommand("copy");
    });
});