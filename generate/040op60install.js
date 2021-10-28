"use strict";

// Initialize deferredPrompt for use later to show browser install prompt.
let deferredPrompt;

window.addEventListener('beforeinstallprompt', (e) => {

    let showInstallPromotion = () => {
	let buttonInstall = document.querySelector("#installeren");
	buttonInstall.classList.remove("hidden040");

	buttonInstall.addEventListener('click', () => {
	    // Hide the app provided install promotion
//	    buttonInstall.classList.add("hidden040");
	    buttonInstall.classList.remove("mdl-button--colored");
	    // Show the install prompt
//	    deferredPrompt.prompt();
	    // We've used the prompt, and can't use it again, throw it away
//	    deferredPrompt = null;
	});
    }

    // Prevent the mini-infobar from appearing on mobile
    e.preventDefault();
    // Stash the event so it can be triggered later.
    deferredPrompt = e;
    // Update UI notify the user they can install the PWA
    showInstallPromotion();
    // Optionally, send analytics event that PWA install promo was shown.
    console.log(`'beforeinstallprompt' event was fired.`);
});

console.log("install event registered");
