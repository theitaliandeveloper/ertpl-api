// Roba da non modificare assolutamente
export const config = {
  runtime: "edge"
};

const corsHeaders = {
  "Access-Control-Allow-Origin": "*",
  "Access-Control-Allow-Methods": "GET, OPTIONS",
  "Access-Control-Allow-Headers": "Content-Type"
};

// Lista dei server modificabile
const servers = [
    {
      name: "Serverissimo",
      url: "https://setaapi.serverissimo.freeddns.org/allnews"
    },
    {
      name: "Vichingo455",
      url: "https://api.vichingo455.qzz.io/seta/allnews"
    },
    {
      name: "Vichingo455 Offsite",
      url: "https://setaapioffsite.vichingo455.qzz.io/allnews"
    }
  ];

// Altra roba da modificare solamente in caso di bisogno
async function checkServer(server, timeoutMs = 800) {
  try {
    const controller = new AbortController();
    const timeout = setTimeout(() => controller.abort(), timeoutMs);
    const response = await fetch(server.url, {
      Range: "bytes=0-0",
      signal: controller.signal
    });
    clearTimeout(timeout);
    return response.ok;
  } catch {
    return false;
  }
}

export default async function handler() {
  for (const server of servers) {
    if (await checkServer(server)) {
      return new Response(
        JSON.stringify({
          status: "ok",
          server: server.name,
          url: server.url.replace("/allnews","")
        }),
        {
          status: 200,
          headers: {
            ...corsHeaders,
            "Content-Type": "application/json",
            "Cache-Control": "public, s-maxage=60, stale-while-revalidate=30"
          }
        }
      );
    }
  }

  return new Response(
    JSON.stringify({ status: "offline" }),
    {
      status: 503,
      headers: {
        ...corsHeaders,
        "Content-Type": "application/json",
        "Cache-Control": "public, s-maxage=30"
      }
    }
  );
}