// Roba da non modificare assolutamente
export const config = {
  runtime: "edge"
};

const corsHeaders = {
  //"Access-Control-Allow-Origin": "https://ertpl.pages.dev",
  "Access-Control-Allow-Origin": "*",
  "Access-Control-Allow-Methods": "GET, OPTIONS",
  "Access-Control-Allow-Headers": "Content-Type"
};

// Lista dei server modificabile
const servers = [
    {
      name: "Serverissimo",
      url: "https://corsesopp.startapi.serverissimo.com/versione"
    },
    {
      name: "DaniLab",
      url: "https://startsopp.daninet.freeddns.org/versione"
    },
    {
      name: "Vichingo455",
      url: "https://api.vichingo455.com/start-corsesopp.json/versione"
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
    if (!response.ok) return false;
    const contentType = response.headers.get("content-type");
    if (!contentType?.includes("application/json")) return false;
    await response.json();
    return true;
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
          url: server.url.replace("/versione","")
        }),
        {
          status: 200,
          headers: {
            ...corsHeaders,
            "Content-Type": "application/json",
            "Cache-Control": "public, s-maxage=30"
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
        "Cache-Control": "no-store"
      }
    }
  );
}