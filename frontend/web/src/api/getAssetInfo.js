export default (baseUrl) => {
    let assetURL = baseUrl + `/assets`
    return fetch(assetURL, {
        method: "GET",
    }).then(resp => {
        console.log("Got response: ", resp)
        return resp.json();
    });
};