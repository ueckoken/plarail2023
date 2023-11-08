import * as DGRAM from 'dgram';
import * as HTTP from 'http';
import * as WS from 'ws';

const
udp=DGRAM.createSocket('udp4'),
svr=HTTP.createServer((req,res)=>(
	res.writeHead(200,{'Content-Type':'text/html'}),
	res.end(`
	<!DOCTYPE html>
	<html lang="en" dir="ltr">
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width,initial-scale=1">
		<title>video</title>
	</head>
	<body>
		<style>:root,body,#img{width:100%;height:100%;margin:0;background-color:#00f;object-fit:contain;vertical-align:top;}</style>
		<img id="img">
		<script>
			'use strict';
			let ws={},t;
			const
				main=_=>(
					ws=Object.assign(new WebSocket(\`ws://\${location.hostname}/ws\`),{
						binaryType:'arraybuffer',
						onopen:_=>console.log('Opened'),
						onclose:_=>console.log('Closed',main()),
						onmessage:e=>img.src='data:image/jpeg;base64,'+btoa(String.fromCharCode(...new Uint8Array(e.data)))
					})
				);
			img.onload=_=>console.log(Math.round(1000/(-t+(t=performance.now()))));
			document.onvisibilitychange=_=>ws.send('');
			onload=main;
		</script>
	</body>
	</html>
	`)
)),
ws=new Set(),
wss=new WS.WebSocketServer({server:svr,path:'/ws'});

wss.on('connection',_=>ws.add(_));
wss.on('close',_=>ws.delete(_));

let fb={},t;
udp.on('message',(x,i)=>(
	x={
		tag:x.subarray(0,4)+'',
		t:x.readUInt32LE(4),
		i:x[8],
		x:x.subarray(9)
	},
	({
		STRT:_=>(fb[x.t]=[...Array(x.i)],setTimeout(_=>delete fb[x.t],500)),
		DATA:_=>fb[x.t]&&(
			fb[x.t][x.i]=x.x,
			fb[x.t].every(_=>_)&&(
				_=Buffer.concat(fb[x.t]),delete fb[x.t],ws.forEach(x=>x.send(_)),
				console.log({fps:Math.round(1000/(-t+(t=Date.now()))),size:_.length,t:x.t})
			)
		),
	}[x.tag])()
	// console.log(x+'',i.address,i.port)
));
udp.on('listening',_=>console.log(`port ${udp.address().port} ...`));
udp.bind(3333);
svr.listen(80);