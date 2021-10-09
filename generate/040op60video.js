"use strict";

// window.addEventListener('DOMContentLoaded', (event) => {
//     console.log('DOM fully loaded and parsed');

(() => {
    let tag = document.createElement('script');
//    tag.id = 'iframe-demo';
    tag.src = 'https://www.youtube.com/iframe_api';
    let firstScriptTag = document.getElementsByTagName('script')[0];
    firstScriptTag.parentNode.insertBefore(tag, firstScriptTag);
})();


let videoElement = document.getElementById('existing-iframe-example');

history.replaceState({"geroteerd": false}, "Zet 'm op 60");
let oldRoot = null;

function _04060roteerVideo(el) {
    if (!el.classList.contains("_040op60rotated")) {
	oldRoot = el.parentNode;
	let body = document.querySelector("body");
	body.insertBefore(el, body.childNodes[0]);
	history.pushState({"geroteerd": true}, "Geroteerde video");
	el.classList.add("_040op60rotated");
    }
}

function _04060normaleVideo(el) {
    if (el.classList.contains("_040op60rotated")) {
	if (oldRoot != null) {
	    oldRoot.appendChild(el);
	}
	el.classList.remove("_040op60rotated")
    }
    if (history.state && history.state.geroteerd) {
	history.go(-1);
    }
}

window.addEventListener('popstate', function(event) {
    console.log("OnPopState", event);
    if (history.state) {
	if (history.state.geroteerd) {
	    _04060roteerVideo(videoElement);
	} else if (history.state.geroteerd === false) {
	    _04060normaleVideo(videoElement);
	}
    }
})

function rotate(playerStatus) {
    let videoElement = document.getElementById('existing-iframe-example');

    if (document.documentElement.clientWidth > document.documentElement.clientHeight) {
	return;
    }

    // playing // paused // buffering
    if (playerStatus == 1 || playerStatus === 2 || playerStatus === 3) {
	_04060roteerVideo(videoElement);
    } else {
	_04060normaleVideo(videoElement);
    }
}

function onPlayerStateChange(event) {
    console.log("playerstatechange");
    rotate(event.data);
}

function onPlayerReady(event) {
    console.log("onplayerready");
    player.addEventListener("onStateChange", onPlayerStateChange);
}

var player;
function onYouTubeIframeAPIReady() {
    console.log("YoutubeIframeAPIReady");
    player = new YT.Player('existing-iframe-example', {
        events: {
	    'onReady' : onPlayerReady,
	    'onStateChange': onPlayerStateChange
        }
    });
    console.log("YoutubeIframeAPIFinished", player, "foo");
}


console.log("End of script");
//});
