import type { NextApiRequest, NextApiResponse } from 'next'

const BACKEND_URL = process.env.BACKEND_URL??"http://127.0.0.1:8080";

export async function getUrlByHash(hash: string) {
    const getResponse = await fetch(`${BACKEND_URL}/url?hash=${encodeURIComponent(hash)}`);
    if (getResponse.status === 404) {
        return { status: getResponse.status, data: {error: "not found"}};
    }
    if (getResponse.status !== 200) {
        return { status: getResponse.status, data: {error: "something went wrong"}};
    }

    const responseData = await getResponse.json();
    return { status: getResponse.status, data: responseData};
}

async function handleGET(req: NextApiRequest, res: NextApiResponse) {
    const { status, data } = await getUrlByHash(req.query.hash as string);
    res.status(status).json(data);
}

async function handlePUT(req: NextApiRequest, res: NextApiResponse) {
    const { url } = req.body;

    const putResponse = await fetch(`${BACKEND_URL}/url`, {
        method: "PUT",
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `url=${encodeURIComponent(url)}`,
    })

    res.status(putResponse.status)
    console.log('res', putResponse.status)
    switch (putResponse.status) {
        case 201:
            const responseData = await putResponse.json()
            res.json(responseData);
            return
        case 400:
            res.json({error: "bad request"})
            return;
        default:
            res.json({error: "something went wrong"});
    }
}

export default function handler(req: NextApiRequest, res: NextApiResponse) {
    switch (req.method) {
        case "PUT":
            handlePUT(req, res);
            return

        case "GET":
            handleGET(req, res);
            return;

        default:
            res.setHeader('Allow', ['GET', 'PUT']);
            res.status(405).end(`Method ${req.method} Not Allowed`);
    }
}
