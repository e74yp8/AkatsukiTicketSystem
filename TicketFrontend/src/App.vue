<script setup>
import {QrcodeStream} from 'vue-qrcode-reader'
import {ref} from "vue";

const result = ref('')
const error = ref('')

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

function onDetect(detectedCodes) {
  result.value = detectedCodes.map(code => code.rawValue)[0]
  fetch('http://localhost:8080/ticket/check/' + result.value)
      .then(response => {
        console.log(response);
        console.log(response.status);
      })

}

</script>

<template>
  <div style="height: 60%">
    <qrcode-stream :track="paintBoundingBox" @detect="onDetect" @error="onError"></qrcode-stream>
  </div>
  <p>Last result: <b>{{ result }}</b></p>
</template>