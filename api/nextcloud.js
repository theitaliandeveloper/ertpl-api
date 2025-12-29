export const config = {
  runtime: "edge"
};

const servers = [
    {
      name: "Serverissimo",
      url: "https://drive.serverissimo.freeddns.org/status.php"
    },
    {
      name: "Vichingo455",
      url: "https://drive.vichingo455.freeddns.org/status.php"
    }
  ];

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
          url: server.url.replace("/status.php", "")
        }),
        {
          status: 200,
          headers: {
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
        "Content-Type": "application/json",
        "Cache-Control": "public, s-maxage=30"
      }
    }
  );
}