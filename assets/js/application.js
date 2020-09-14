require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");


function copy() {
    var copyText = document.querySelector("#bpstring")
    copyText.select()
    console.log("Copied", copyText.value)
    document.execCommand("copy");
}

$(() => {
    document.querySelector("#copy").addEventListener("click", copy);
});
