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

            document.querySelector("#result1").value = data.result1
            document.querySelector("#result2").value = data.result2
        })

}
