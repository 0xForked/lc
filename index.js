const fs = require('fs')
const path = require('path')

const args = process.argv
    .slice(2)
    .toString()
    .split(',')

const helpFlags = ["-h",  "--h", "--help", "help"]
const fileFlags = ["-f", "--f", "--file", "file"]
const extensionFlags =  ["-t", "--t", "--transform", "transform"]
const outputDirFlags = ["-o", "--o", "--output", "output"]

let filePath = null
let transformExtension = "json"
let targetDirectory = "./storage"

if (args.length > 0) {
    if (helpFlags.includes(args[0])) {
        help()
        return
    }

    args.forEach((element, index) => {
        if (fileFlags.includes(element)) {
            filePath = args[index + 1]
        }

        if (extensionFlags.includes(element)) {
            transformExtension = args[index + 1]
        }

        if (outputDirFlags.includes(element)) {
            targetDirectory = args[index + 1]
        }
    });

    if (filePath != null) {
        transformData(filePath, transformExtension, targetDirectory)
        return
    }

    home()
}

function home() {
   console.log("display home")
}

function help() {
    console.log("Display Options")
}

function transformData(target, extension, output) {
    let logData = []

    const filePath = path.join(target)
    const fileData = fs.createReadStream(filePath, 'utf8');

    fileData.on("data", (data) => {
        const splitData = data.split("\n")

        splitData.forEach(element => {
            if (extension === "json") {
                logData.push({"data": element})
            }

            if (extension === "text") {
                logData.push(element)
            }
        })
    })

    fileData.on('end', () => {
        if (extension === "json") {
            const data = JSON.stringify(logData, null, 2)
            fs.writeFileSync(`${output}/data.json`, data)
            console.log("json file crated")
            logData = []
            return
        }

        if (extension === "text") {
            logData.forEach(data => {
                fs.appendFile(`${output}/data.txt`, `${data}\r\n`, function (err) {
                    if (err) throw err;
                });
            })
            console.log("txt file crated")
            logData = []
            return
        }

        console.log("failed convert file")
    })
}