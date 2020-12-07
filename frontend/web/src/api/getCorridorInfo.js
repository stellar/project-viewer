export default (baseUrl, fromCode, fromIssuer, toCode, toIssuer) => {
    let corridorURL = baseUrl + `/corridor?sourceCode=${fromCode}&sourceIssuer=${fromIssuer}&destCode=${toCode}&destIssuer=${toIssuer}`
    return fetch(corridorURL, {
        method: "GET",
    }).then(resp => {
        console.log("Got response: ", resp)
        return resp.json();
    });
};