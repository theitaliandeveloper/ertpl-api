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
      url: "https://startsopp.serverissimo.freeddns.org/"
    },
    {
      name: "Vichingo455",
      url: "https://api.vichingo455.qzz.io/start-corsesopp.json"
    },
    {
      name: "Vichingo455 Offsite",
      url: "https://startcorsesoppoffsite.vichingo455.qzz.io/"
    }
  ];

// Altra roba da modificare solamente in caso di bisogno
async function checkServer(server, timeoutMs = 1500) {
  try {
    const controller = new AbortController();
    const timeout = setTimeout(() => controller.abort(), timeoutMs);
    const response = await fetch(server.url, {
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
          url: server.url
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