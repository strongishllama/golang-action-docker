const process = require("process");
const childProcess = require("child_process");

function debugLog(message) {
  if (process.env["ACTIONS_STEP_DEBUG"]) {
    console.debug(message);
  }
}

function getBinaryPath() {
  let binaryPath = `${__dirname}/bin/action-${process.platform}-`;

  if (process.platform === "linux" && process.arch === "x64") {
    binaryPath += "amd64";
  } else {
    console.error(`Unsupported platform: ${process.platform} or architecture: ${process.arch}`);
    process.exit(1);
  }

  return binaryPath;
}

const binaryPath = getBinaryPath();
debugLog(`Executing binary at: ${binaryPath}`);

const spawnSyncReturns = childProcess.spawnSync(binaryPath, { stdio: "inherit" });
debugLog(JSON.stringify(spawnSyncReturns));

if (spawnSyncReturns.status !== 0) {
  process.exit(1);
}
