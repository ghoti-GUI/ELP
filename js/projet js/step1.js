
function Process(){
    const { execSync } = require('child_process');
    const output = execSync('ps').toString();
    console.log(output);
}
function fondshellprogram(path){
    const { spawn } = require('child_process');
    const program = spawn('bash',[path]);

    program.stdout.on('data', (data) => {
        console.log(` ${data}`);
    });

    program.stderr.on('data', (data) => {
        console.log(`error: ${data}`);
    });

    program.on('close', (code) => {
        console.log(`shell progress exited with code ${code}`);
    });
}
function killprocess(pid){
    const { execSync } = require('child_process');

    try {
        execSync(`kill -9 ${pid}`);
        console.log(`Process with pid ${pid} has been terminated successfully.`);
    } catch (error) {
        console.error(`Failed to terminate the process: ${error}`);
    }
        console.log(`process ${pid} is killed`)
}

function stopprocess(pid){
    const { execSync } = require('child_process');
    /*execSync(`kill -STOP ${pid}`);*/
    const output = execSync(`kill -STOP ${pid}`).toString();
    console.log(output);
}

function restartprocess(pid){
    const { execSync } = require('child_process');
    /*execSync(`kill -CONT ${pid}`);*/
    const output = execSync(`kill -CONT ${pid}`).toString();
    console.log(output);
}

function shellprogram(path){
    const { spawnSync } = require('child_process');
    const result = spawnSync('bash',[path]);
    console.log(result.stdout.toString());    
}

function jobs(){
    const { execSync } = require('child_process');
    const output = execSync('jobs').toString();
    console.log(output);
}

function keepprocess(pid){
    /*const { execSync } = require('child_process');
    const output = execSync(`bg`).toString();
    console.log(output);*/
    const { exec } = require('child_process');
    exec(`kill -CONT ${pid} `,(err,stdout,stderr) => {
        if (err) {
            console.error(err);
            return
        }
        console.log(stdout)
        console.log(stderr)
    })
}
function backlinux(input){
    const { exec } = require('child_process');
    exec(`${input} `,(err,stdout,stderr) => {
        if (err) {
            console.error(err);
            return
        }
        console.log(stdout)
        console.log(stderr)
    })
}
function linux(input){
    const { execSync } = require('child_process');

try {
  const stdout = execSync(`${input}`).toString();
  console.log(stdout);
} catch (error) {
  console.error(error);
}
}
function backlinux2(input){
    const { spawn } = require('child_process');
    const program = spawn(`${input}`);

    program.stdout.on('data', (data) => {
        console.log(` ${data}`);
    });

    program.stderr.on('data', (data) => {
        console.log(`error: ${data}`);
    });

    program.on('close', (code) => {
        console.log(`shell code ${code}`);
    });
}
const readline = require('readline');
const r1 = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});
readline.emitKeypressEvents(process.stdin)
process.stdin.setRawMode(true);
process.stdin.on('keypress', (str,key) => {
    if (key.ctrl&&key.name==='p') {
        console.log("exit with control p")
        process.exit();
    }
});
readUserInput();

function readUserInput(){
r1.question("command: ", (input) => {

if(input == 'exit') {
    r1.close();
}else{
    let order = input.split(' ');
    switch(order[0]){
        case 'lp': 
            Process();
            break
        case 'bing':
            if(order[1]=='-k'){
                console.log("start to kill")
                killprocess(order[2]);
            }else{
                if(order[1]=='-p'){
                    stopprocess(order[2]);
                }else{
                    if(order[1]=='-c'){
                        restartprocess(order[2]);
                    }else{
                        console.log("wrong command");
                    }
                }
            }
            break;
        case 'keep':
            stopprocess(order[1]);
            keepprocess(order[1]);
            break;
        case 'jobs':
            console.log('jobs')
            jobs();
            break;
        default:
            if(input == ''||input==' '){
                //console.log('rien')
                break;
            }else{
                if (input.endsWith('.sh')||input.endsWith('.sh !')){
                    if(order.length==1){
                        //console.log('shellbegin')
                        shellprogram(order[0]);
                        break;
                    }else{
                        
                        //console.log('shellbegin')
                        fondshellprogram(order[0]);
                        break;
                        
                    }
                }else{
                    if (input.endsWith('!')){
                        backlinux2(input.slice(0, -2));
                        break;
                    }else{
                        linux(input);
                        break;
                    }
                }
            }
                
            break
    }
    readUserInput()
}
})
}
