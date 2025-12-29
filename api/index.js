export const config = {
  runtime: "edge"
};

export default function handler() {
  return new Response(
    JSON.stringify({
      name: "ER-TPL API",
      version: "1.0.0",
      description: "API to check NextCloud servers status for ERTPL",
      endpoints: {
        nextcloud: "/nextcloud"
      }
    }),
    {
      headers: {
        "Content-Type": "application/json",
        "Cache-Control": "public, s-maxage=300"
      }
    }
  );
}