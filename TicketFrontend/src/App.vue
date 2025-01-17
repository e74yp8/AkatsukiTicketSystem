<script setup>
import {QrcodeStream} from 'vue-qrcode-reader'
import {ref} from "vue";
import axios from "axios";


const result = ref('')
const error = ref('')
const id = ref('')
const status = ref('')
let bgColor = '#ffffff'

function paintBoundingBox(detectedCodes, ctx) {
  for (const detectedCode of detectedCodes) {
    const {
      boundingBox: {x, y, width, height}
    } = detectedCode

    ctx.lineWidth = 2
    ctx.strokeStyle = '#007bff'
    ctx.strokeRect(x, y, width, height)
  }
}

function onError(err) {
  error.value = `[${err.name}]: `

  if (err.name === 'NotAllowedError') {
    error.value += 'you need to grant camera access permission'
  } else if (err.name === 'NotFoundError') {
    error.value += 'no camera on this device'
  } else if (err.name === 'NotSupportedError') {
    error.value += 'secure context required (HTTPS, localhost)'
  } else if (err.name === 'NotReadableError') {
    error.value += 'is the camera already in use?'
  } else if (err.name === 'OverconstrainedError') {
    error.value += 'installed cameras are not suitable'
  } else if (err.name === 'StreamApiNotSupportedError') {
    error.value += 'Stream API is not supported in this browser'
  } else if (err.name === 'InsecureContextError') {
    error.value += 'Camera access is only permitted in secure context. Use HTTPS or localhost rather than HTTP.'
  } else {
    error.value += err.message
  }
}

async function onDetect(detectedCodes) {
  result.value = detectedCodes.map(code => code.rawValue)[0]

  const URL = "/api/ticket/check/" + result.value;
  try {
    axios.get(URL, {
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => {
      id.value = response.data.id;
      status.value = response.data.status;
      if (status.value === "可入場") {
        bgColor = "#69e62b"
      } else {
        bgColor = "#fa5353"
      }
      document.body.style.backgroundColor = bgColor;
    }).catch(error => {
      console.error('發生錯誤！', error);
      error.value = error;
    });
  } catch (e) {
    console.log(e);
    error.value = e;
  }
}

function reset_data() {
  id.value = "";
  status.value = "";
  error.value = "";
  document.body.style.backgroundColor = "#ffffff";
}

</script>


<style>
html, body {
  height: 100%;
  margin: 0;
  font-family: Arial, sans-serif;
  display: flex;
  align-items: center;
}

.qrcode-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  background: rgba(0, 0, 0, 0.1);
}

.qrcode-stream {
  border: 2px solid #007bff;
  border-radius: 10px;
  overflow: hidden;
  max-width: 600px;
}

.id {
  margin-top: 20px;
  font-size: 1.2em;
  color: #333;
}

.error-message {
  color: red;
}

#app {
  height: 60%;
}

h1 {
  position: fixed;
  top: 20px;
  left: 50%;
  transform: translateX(-50%);
  white-space: nowrap;
}

button {
  margin-top: 20px;
  padding: 10px 20px;
  font-size: 1em;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

</style>

<template>
  <h1>{{ status }}</h1>
  <div class="qrcode-container">
    <qrcode-stream class="qrcode-stream" :track="paintBoundingBox" @detect="onDetect" @error="onError"></qrcode-stream>
    <p class="id">籌號: <b>{{ id }}</b></p>
    <p class="error-message" v-if="error">{{ error }}</p>
  </div>

  <div style="width: 100%; display: flex; justify-content: center;">
    <button @click="reset_data">重設</button>
  </div>
</template>
