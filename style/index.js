// alert("hello")

let data = document.querySelector(".myForm")
document.querySelector(".submit").addEventListener("click", req)

// document.querySelector("#result1").value = 5
function req(e) {
    console.log("i got triggered")

    e.preventDefault()
    let dataForm = new FormData(data)

    fetch("/Result", {
        method: "post",
        body: dataForm,
    })
        .then(res => res.json())
        .then((data) => {

            if (data.Error !== "") {
                document.querySelector(".error-code").innerText = "Error " + "400";
                document.querySelector(".error-message").innerText = data.Error;
                document.querySelector("#prevent1").style.display = "block";
                document.querySelector("#dialog").style.display = "flex";
                return
            }

            document.querySelector("#result1").value = data.Result1
            document.querySelector("#result2").value = data.Result2
        })

}


document.querySelector(".close-button").addEventListener("click", () => {
    document.querySelector("#dialog").style.display = "none";
    document.querySelector("#prevent1").style.display = "none";
})