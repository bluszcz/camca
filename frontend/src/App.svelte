<script lang="ts">
  import logo from "./assets/images/logo-universal.png";
  import { CalcSensorRatio, Greet } from "../wailsjs/go/main/App.js";
  import { CalcPixDens } from "../wailsjs/go/main/App.js";
  let resultPDX: string = "μm";
  let resultSensorRatio: string = "";
  let name: string;
  let sensorX: string;
  let sensorY: string;
  let pixelWidth: string;
  let FOV: string;
  let resultFOV: string;
  let cropFactor: string;
  let fps: string;

  let cameras = [
    { id: 0, name: "Custom", sensorX: "", sensorY: "", pixelWidth: "" },
    {
      id: 1,
      name: "BMPCC4K",
      sensorX: "18.96",
      sensorY: "10",
      pixelWidth: "4096",
    },
    {
      id: 2,
      name: "Lumix D9M2",
      sensorX: "17.3",
      sensorY: "13",
      pixelWidth: "5760",
    },
    {
      id: 3,
      name: "Canon R7",
      sensorX: "22.2",
      sensorY: "14.8",
      pixelWidth: "6960",
    },
    {
      id: 4,
      name: "Canon M50 M2",
      sensorX: "22.3",
      sensorY: "14.9",
      pixelWidth: "6000",
    },
  ];
  let cameraSelected;
  let answer = "";

  const onCameraChange = () => {
    sensorX = cameraSelected.sensorX;
    sensorY = cameraSelected.sensorY;
    pixelWidth = cameraSelected.pixelWidth;
    calculateCamera();
  };

  const onSensorX = () => {
    cameraSelected = cameras[0];
    calculateCamera();
  };
  const onSensorY = () => {
    cameraSelected = cameras[0];
    calculateCamera();
  };
  const onPixelWIdth = () => {
    cameraSelected = cameras[0];
    calculateCamera();
  };

  const onFOV = () => {
    // cameraSelected = cameras[0];
    // calculateCamera();
  };


  const onCropFactor = () => {
    // cameraSelected = cameras[0];
    // calculateCamera();
  };

  function calculateCamera(): void {
    calcuatePD();
    calcuateSensorRatio();
  }

  function calcuatePD(): void {
    CalcPixDens(sensorX, sensorY, pixelWidth).then(
      (result) => (resultPDX = result + " μm"),
    );
  }



  function calcuateSensorRatio(): void {
    CalcSensorRatio(sensorX, sensorY).then(
      (result) => (resultSensorRatio = result),
    );
  }

  import { onMount } from "svelte";

  // let canvas;

  // onMount(() => {
  //   const ctx = canvas.getContext("2d");
  //   let frame;

  //   (function loop() {
  //     frame = requestAnimationFrame(loop);

  //     const imageData = ctx.getImageData(0, 0, canvas.width, canvas.height);

  //     for (let p = 0; p < imageData.data.length; p += 4) {
  //       const i = p / 4;
  //       const x = i % canvas.width;
  //       const y = (i / canvas.height) >>> 0;

  //       const t = window.performance.now();

  //       const r = 64 + (128 * x) / canvas.width + 64 * Math.sin(t / 1000);
  //       const g = 64 + (128 * y) / canvas.height + 64 * Math.cos(t / 1400);
  //       const b = 128;

  //       imageData.data[p + 0] = r;
  //       imageData.data[p + 1] = g;
  //       imageData.data[p + 2] = b;
  //       imageData.data[p + 3] = 255;
  //     }

  //     ctx.putImageData(imageData, 0, 0);
  //   })();

  //   return () => {
  //     cancelAnimationFrame(frame);
  //   };
  // });
</script>

<main>
  <img alt="Camca logo" id="logo" src={logo} />
  <div class="grid-container">
    <div class="grid-item">
      <div class="inputs">
        <div>Pixel Density</div>

        <hr />
        Camera:<select bind:value={cameraSelected} on:change={onCameraChange}>
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
            bind:value={sensorX}
            on:change={onSensorX}
            class="input"
            id="sensorX"
            type="text"
          />
        </div>
        <div class="input-box">
          <label for="sensorX">Sensor height mm</label>
          <input
            autocomplete="off"
            bind:value={sensorY}
            on:change={onSensorY}
            class="input"
            id="sensorY"
            type="text"
          />
        </div>
        <div class="input-box">
          <label for="widthPx">Width px</label>

          <input
            autocomplete="off"
            bind:value={pixelWidth}
            on:change={onPixelWIdth}
            class="input"
            id="pixelWidth"
            type="text"
          />
        </div>
        <hr />
        <div>Results</div>
        <table class="results">
          <tr>
            <td>Pixel density</td>
            <td>{resultPDX}</td>
          </tr>
          <tr>
            <td>Sensor ratio</td>
            <td>{resultSensorRatio}</td>
          </tr>
        </table>
        <hr />
      </div>
    </div>
    <div class="grid-item">
      <div class="inputs">
        Field of View Crop Factor
        <hr/>
        <div class="input-box">
          <label for="FOV">Field of View mm:</label>
          <input
            autocomplete="off"
            bind:value={FOV}
            on:change={onFOV}
            class="input"
            id="FOV"
            type="text"
          />
        </div>
        <div class="input-box">
          <label for="cropFactor">Crop Factor</label>
          <input
            autocomplete="off"
            bind:value={cropFactor}
            on:change={onCropFactor}
            class="input"
            id="cropFactor"
            type="text"
          />
        </div>
        <hr/>
        <div>Results</div>
        <table class="results">
          <tr>
            <td>New FOV:</td>
            <td>{resultFOV}</td>
          </tr>
          <!-- <tr>
            <td>Sensor ratio</td>
            <td>{resultSensorRatio}</td>
          </tr>
        </table> -->
      </table> 
        <hr />

      </div>
    </div>
    <div class="grid-item">
      <div class="inputs">
        180 deegre Shutter Speed
        <hr/>
        <div class="input-box">
          <label for="fps">Frames per second fps:</label>
          <input
            autocomplete="off"
            bind:value={fps}
            on:change={onfps}
            class="input"
            id="fps"
            type="text"
          />
        </div>
        <div class="input-box">
          <label for="cropFactor">Angle </label>
          <input
            autocomplete="off"
            bind:value={cropFactor}
            on:change={onCropFactor}
            class="input"
            id="cropFactor"
            type="text"
          />
        </div>
        <div class="input-box">
          <label for="cropFactor">Shutter Speed 1/</label>
          <input
            autocomplete="off"
            bind:value={cropFactor}
            on:change={onCropFactor}
            class="input"
            id="cropFactor"
            type="text"
          />
        </div>
        <hr/>


      </div>
    </div>
  </div>
  <!-- FOV Calculator -->
  <!-- 180 Shutter Rule Video Calculator -->
  <!-- <button class="btn" on:click={calcuatePD}>Calculates</button> -->
  <!-- CANVAS GRAPH AREA -->
  <!-- <div>
  <canvas bind:this={canvas} width={320} height={320} />
</div> -->
</main>

<style>
  .grid-container {
  display: grid;
   grid-template-columns: auto auto auto;
   padding: 10px;


}
table {
  width: 600px;
}

.grid-item {
  
}
  td:nth-child(1) {
    text-align: left;
  }
  td:nth-child(2) {
    text-align: right;
  }
  table.results {
    width: 200px;
    margin: 0 auto;
  }
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
    width: 300px;
    margin: 0 auto;
    /* text-align: center; */
  }
  label {
    display: inline-block;
    width: 200px;
    text-align: left;
  }

  .input-box .div {
    /* text-align: right; */
    margin-bottom: 10px;
    width: 100%;
  }

  .input-box .btn {
    width: 80px;
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
    width: 60px;
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
