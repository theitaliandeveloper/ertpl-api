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
      url: "https://drive.serverissimo.freeddns.org/status.php",
      share: "ffdqobqRg2ezKXt"
    },
    {
      name: "Vichingo455",
      url: "https://drive.vichingo455.qzz.io/status.php",
      share: "w8Nr4jZN3g6z3pn"
    },
    {
      name: "Vichingo455 Offsite",
      url: "https://driveoffsite.vichingo455.qzz.io/status.php",
      share: "b6W9NLF7mpDwZQE"
    }
  ];

// Altra roba da modificare solamente in caso di bisogno
async function checkServer(server, timeoutMs = 800) {
  try {
    const controller = new AbortController();
    const timeout = setTimeout(() => controller.abort(), timeoutMs);

    const response = await fetch(server.url, {
      signal: controller.signal
    });

    clearTimeout(timeout);

    if (!response.ok) return false;

    const data = await response.json();
    return data.installed === true && data.maintenance === false;
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
          url: server.url.replace("/status.php", ""),
	  share: server.share
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