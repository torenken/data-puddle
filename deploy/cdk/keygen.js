const crypto = require('crypto');
const s = crypto.randomBytes(32).toString("base64");
console.log('encryption key for using 256-bit AES-GCM (base64 encoding)')
console.log('encryption key:', s)