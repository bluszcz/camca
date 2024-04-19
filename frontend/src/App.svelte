<script lang="ts">
  import logo from "./assets/images/logo-universal.png";
  import { Greet } from "../wailsjs/go/main/App.js";
import { CalcPixDens } from "../wailsjs/go/main/App.js";
  let resultPDX: string = "Pixel density:";
  let name: string;
  let sensorX: string;
  let sensorY: string;
  let pixelWidth: string;

	let cameras = [
		{ id: 0, name : "Custom", sensorX:"",sensorY:"", pixelWidth:""},
    
		{ id: 1, name : "BMPCC4K", sensorX:"17.2",sensorY:"10", pixelWidth:"4096"},
		{ id: 2, name : "Lumix D9M2", sensorX:"17.3",sensorY:"13", pixelWidth:"5760"},
		{ id: 3, name : "Canon R7", sensorX:"22.2",sensorY:"14.8", pixelWidth:"6960"},
		{ id: 4, name : "Canon M50 M2", sensorX:"22.3",sensorY:"14.9", pixelWidth:"6000"},
	];
	let cameraSelected;
	let answer = '';

  const onCameraChange = () => {
    sensorX = cameraSelected.sensorX
    sensorY = cameraSelected.sensorY
    pixelWidth = cameraSelected.pixelWidth
    calcuatePD()
  }

  const onSensorX = () => {
    cameraSelected = cameras[0]
    calcuatePD()

  }
  const onSensorY = () => {
    cameraSelected = cameras[0]
    calcuatePD()

  }
  const onPixelWIdth = () => {
    cameraSelected = cameras[0]
    calcuatePD()

  }


  function calcuatePD(): void {
    CalcPixDens(sensorX, sensorY, pixelWidth).then((result) => (resultPDX = "Pixel Density: "+ result+" μm"));
  }
</script>

<main>
  <img alt="Camca logo" id="logo" src={logo} />
  <div class="result" id="result">{resultPDX}</div>
  <div class="inputs">

    <select bind:value={cameraSelected} on:change={onCameraChange}>
      {#each cameras as camera}
        <option value={camera}>
          {camera.name}
        </option>
      {/each}
    </select>



    <div class="input-box">
      <label for="sensorX">Sensor width mm</label>
      <input
        autocomplete="off"
        bind:value={sensorX} on:change={onSensorX}
        class="input"
        id="sensorX"
        type="text"
      />
    </div>

    <div class="input-box">
      <label for="sensorX">Sensor height mm</label>

      <input
        autocomplete="off"
        bind:value={sensorY}  on:change={onSensorY}
        class="input"
        id="sensorY"
        type="text"
      />
    </div>

    <div class="input-box">
      <label for="widthPx">Width px</label>

      <input
        autocomplete="off"
        bind:value={pixelWidth} on:change={onPixelWIdth}
        class="input"
        id="pixelWidth"
        type="text"
      />
    </div>
  </div>
  <button class="btn" on:click={calcuatePD}>Calculates</button>
</main>

<style>
  #logo {
    display: block;
    width: 20%;
    height: 20%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .inputs {
    width: 600px;
    margin: 0 auto;
    text-align: center;

  }
  label {
    display: inline-block;
            width: 30%; 
  }

  .input-box .div {
    text-align: right;
            margin-bottom: 10px; 
  }

  .input-box .btn {
    width: 160px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;

  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style>
