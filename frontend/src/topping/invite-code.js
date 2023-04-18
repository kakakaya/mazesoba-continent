export function GenerateDummyInviteCode() {
    return "bsky-social-"+generateRandomString(7);
}

function generateRandomString(length) {
    const characters = 'abcdefghijklmnopqrstuvwxyz0123456789';
    let result = '';

    while (true) {
        if (result.length === length) {
            const regex = /\d/g;
            const matches = result.match(regex);
            if (matches && matches.length > 1) {
                console.log(result);
                result = '';    // reset
            } else {
                return result;   // OK
            }
        }
        result += characters.charAt(Math.floor(Math.random() * characters.length));
    }
}

