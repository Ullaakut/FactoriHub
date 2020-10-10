require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

$(() => {
    let blueprints = document.getElementsByClassName("copy")

    for (var i=0; i < blueprints.length; i++) {
        blueprints[i].onclick = event => {
            let id = event.target.id
            if (id === "") {
                id = event.target.parentNode.id
            }
            var copyText = document.getElementById(id).childNodes[3]
            console.log(copyText)
            copyText.select()
            console.log("Copied", copyText.value)
            document.execCommand("copy");
        };
    }
});
