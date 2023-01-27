/*
function Process(){
    const { exec } = require('child_process');
    exec('tasklist', (error, stdout, stderr) => {
        if (error) {
            console.error(`${error}`);
            return;
        }
        console.log( `${stdout}`);
    });
}
*/

function Process(){
    const { execSync } = require('child_process');
    const output = execSync('tasklist').toString();
    //console.log(output);
    return output;
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
    const { spawn } = require('child_process');
    const program = spawn('taskkill',[pid]);
    program.stdout.on('data', (data) => {
        console.log(` ${data}`);
    });

    program.stderr.on('data', (data) => {
        console.log(`error: ${data}`);
    });

    program.on('close', (code) => {
        console.log(`exited with code ${code}`);
    });
}


function shellprogram(path){
    const { spawnSync } = require('child_process');
    const result = spawnSync('bash',[path]);
    console.log(result.stdout.toString());
    console.log(result.stderr.toString());
}







const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
  });

readline.on('SIGINT', () => {
    console.log("You pressed Ctrl + P");
    readline.close();
});

function main_function(){
    readline.question("Enter some words separated by space (Ctrl + P to quit) : ", (input) => {
        let order = input.split(" ");
        switch(order[0]){
            case 'lk': 
                output = Process();
                console.log(output);
                console.log("end..........");
                break;
            case 'bing':
                if(order[2]=='-k'){
                    killprocess(order[3]);
                }
                break;
            default:
                if(order.length==1){
                    shellprogram(order[0]);
                }else{
                    if(length(order)==2&&order[2]=='!'){
                        fondshellprogram(order[0]);
                    }else{
                        console.log(`error\n`);
                    }
                }
                break;
        }
        console.log(order);
        main_function();
    })
}


main_function()