
var child_process = require('child_process')

child_process.execSync('rm -rf basic && bsb -init basic && cd basic && npm run build',{stdio:[0,1,2]})