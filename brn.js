#!/usr/bin/env node
import("./index.js").catch((err) => {
  console.error("Error loading the CLI:", err);
  process.exit(1);
});
