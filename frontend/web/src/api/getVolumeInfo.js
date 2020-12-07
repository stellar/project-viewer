export default (baseUrl, code, issuer, volumeFrom) => {
    let volumeURL = baseUrl + `/volume?code=${code}&issuer=${issuer}&volumeFrom=${volumeFrom}`
    return fetch(volumeURL, {
        method: "GET",
    }).then(resp => {
        console.log("Got response: ", resp)
        return resp.json();
    });
};