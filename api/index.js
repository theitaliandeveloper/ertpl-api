// Roba da modificare solo in caso di bisogno

export const config = {
  runtime: "edge"
};

const corsHeaders = {
  "Access-Control-Allow-Origin": "*",
  "Access-Control-Allow-Methods": "GET, OPTIONS",
  "Access-Control-Allow-Headers": "Content-Type"
};

import pkg from "../package.json";

export default function handler() {
  return new Response(
    JSON.stringify({
      name: pkg.name,
      version: pkg.version,
      description: pkg.description,
      homepage: pkg.homepage
    }),
    {
      headers: {
        ...corsHeaders,
        "Content-Type": "application/json",
        "Cache-Control": "public, s-maxage=300"
      }
    }
  );
}