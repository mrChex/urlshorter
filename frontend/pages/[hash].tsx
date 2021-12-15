import { useEffect } from "react";
import { useRouter } from 'next/router';
import {getUrlByHash} from "./api/url";
import {NextPageContext} from "next";


interface HashPageProps {
    status: number
    data: {
        error?: string
        url?: string    // never provided to page, because of 301 redirect on server side
    }
}

const HashPage = ({data}: HashPageProps) => {
    return <>Error: {data.error}</>
}
export default HashPage

export const getServerSideProps = async (ctx: NextPageContext) => {
    const urlResponse = await getUrlByHash(ctx.query.hash as string) as HashPageProps;
    if (urlResponse.status === 200) {
        console.log('url', urlResponse.data.url)
        ctx.res.setHeader('Location', urlResponse.data.url);
        ctx.res.statusCode = 301;
    }

    return { props: urlResponse }
}