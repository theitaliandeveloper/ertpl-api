export const config = {
  runtime: "edge"
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
        "Content-Type": "application/json",
        "Cache-Control": "public, s-maxage=300"
      }
    }
  );
}