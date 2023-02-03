
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
        case 'lk': 
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
            jobs();
            break;
        default:
            if(order.length==1){
                shellprogram(order[0]);
            }else{
                if(order.length==2&&order[1]=='!'){
                    fondshellprogram(order[0]);
                }else{
                    console.log(`wrong command`);
                }
            }
            break
    }
    
    //console.log(order);
    readUserInput()
}
   
})
}
