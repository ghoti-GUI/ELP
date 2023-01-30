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
	try{
		execSync(`kill -STOP ${pid}`);
        } catch (error) {
                console.error(`Failed to stop the process: ${error}`);
        }
                console.log(`process ${pid} is stopped`)

}
function restartprocess(pid){
	const { execSync } = require('child_process');
	try{
		execSync(`kill -CONT ${pid}`);
	}catch (error){
                console.error(`Failed to restart the process: ${error}`);
        }
                console.log(`process ${pid} is restarteded`)
}
function shellprogram(path){
  console.log('start')
	const { spawnSync } = require('child_process');
	const result = spawnSync('bash',[path]);
	console.log(result.stdout.toString());
	console.log(result.stderr.toString());
	console.log('end')
}
function keepprocess(pid){
	const { execSync } = require('child_process');
	try{
	  const output = execSync(`bg`).toString();
	  console.log(output);
	}catch (error){
	  console.error(`Failed to keep the process: ${error}`);
	}
	console.log(`process ${pid} is kept`)
}


const readline = require('readline').createInterface({
	input: process.stdin,
	output: process.stdout
});

process.stdin.setRawMode(true);
process.stdin.resume();
process.stdin.setEncoding("utf8");
process.stdin.on("data", (key) => {
	if (key == "\u0005") {
		readline.close();
		process.exit();
	}
});

readUserInput();
function readUserInput(){
readline.question("Enter some words separated by space: ", (input) => {
	
if(input == 'exit') {
	readline.close();
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
						console.log("wrong")
					}
				}
			}
			break;
		case 'keep':
		  stopprocess(order[2]);
			keepprocess(order[2]);
			break;
		default:
			if(order.length==1){
				shellprogram(order[0]);
			}else{
				if(order.length==2&&order[1]=='!'){
					fondshellprogram(order[0]);
				}else{
					console.log(`error\n`);
				}
			}
			break
	}
	console.log(order);
	readUserInput()
}
   
})
}
