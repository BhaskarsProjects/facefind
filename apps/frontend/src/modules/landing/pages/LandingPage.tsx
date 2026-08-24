import { Link } from "react-router-dom";
import { useAuth0 } from "@auth0/auth0-react";

import { Button } from "@/shared/components/ui/button";

export default function LandingPage() {
  const { loginWithRedirect } = useAuth0();

  const handleLogin = () => {
    loginWithRedirect();
  };

  const handleSignup = () => {
    loginWithRedirect({
      authorizationParams: {
        screen_hint: "signup",
      },
    });
  };

  const { isAuthenticated, isLoading, user } =
    useAuth0();

  console.log({
    isAuthenticated,
    isLoading,
    user,
  });

  return (
    <main className="min-h-screen bg-background text-foreground">
      {/* Header */}
      <header className="border-b">
        <div className="mx-auto flex h-16 max-w-7xl items-center justify-between px-6">
          <Link to="/" className="text-xl font-bold tracking-tight">
            FaceFind
          </Link>

          <nav className="flex items-center gap-3">
            <Button
              variant="ghost"
              className="px-4 py-2 text-sm font-medium"
              onClick={handleLogin}
            >
              Log in
            </Button>

            <Button
              variant="default"
              className="h-11 px-6 text-sm font-semibold"
              onClick={handleSignup}
            >
              Create account
            </Button>
          </nav>
        </div>
      </header>

      {/* Hero */}
      <section className="mx-auto flex min-h-[calc(100vh-4rem)] max-w-7xl items-center px-6 py-20">
        <div className="grid w-full gap-16 lg:grid-cols-2 lg:items-center">
          {/* Content */}
          <div className="max-w-2xl">
            <div className="mb-6 inline-flex rounded-full border bg-muted/50 px-3 py-1 text-sm text-muted-foreground">
              Find your photos with your face
            </div>

            <h1 className="text-5xl font-bold tracking-tight sm:text-6xl">
              Find every photo of <span className="text-primary">you.</span>
            </h1>

            <p className="mt-6 max-w-xl text-lg leading-8 text-muted-foreground">
              FaceFind uses face recognition to help you quickly discover photos
              of yourself across your entire photo collection.
            </p>

            <div className="mt-8 flex flex-col gap-3 sm:flex-row">
              <Button
                variant="default"
                className="h-11 px-6 text-sm font-semibold"
                onClick={handleSignup}
              >
                Create account
              </Button>

              <Button
                variant="outline"
                className="h-11 px-6 text-sm font-semibold"
                onClick={handleLogin}
              >
                Log in
              </Button>
            </div>

            <p className="mt-4 text-sm text-muted-foreground">
              Simple. Private. Built around your photos.
            </p>
          </div>

          {/* Visual */}
          <div className="relative">
            <div className="aspect-square overflow-hidden rounded-3xl border bg-muted/30 shadow-2xl">
              <div className="flex h-full items-center justify-center p-10">
                <div className="w-full max-w-sm">
                  <div className="rounded-2xl border bg-background p-6 shadow-lg">
                    <div className="flex items-center gap-4">
                      <div className="flex h-16 w-16 items-center justify-center rounded-full bg-primary/10 text-2xl">
                        👤
                      </div>

                      <div>
                        <p className="font-semibold">Your face</p>
                        <p className="text-sm text-muted-foreground">
                          Face detected
                        </p>
                      </div>
                    </div>

                    <div className="mt-6 grid grid-cols-3 gap-2">
                      <div className="aspect-square rounded-lg bg-muted" />
                      <div className="aspect-square rounded-lg bg-muted" />
                      <div className="aspect-square rounded-lg bg-muted" />
                      <div className="aspect-square rounded-lg bg-muted" />
                      <div className="aspect-square rounded-lg bg-muted" />
                      <div className="aspect-square rounded-lg bg-muted" />
                    </div>

                    <div className="mt-6 flex items-center justify-between rounded-lg bg-muted/60 px-4 py-3">
                      <span className="text-sm font-medium">
                        Matching photos
                      </span>

                      <span className="text-sm font-semibold">128 found</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Features */}
      <section className="border-t bg-muted/20">
        <div className="mx-auto max-w-7xl px-6 py-20">
          <div className="max-w-2xl">
            <h2 className="text-3xl font-bold tracking-tight">
              Your photos, easier to find.
            </h2>

            <p className="mt-4 text-muted-foreground">
              FaceFind is designed to make searching through large photo
              collections effortless.
            </p>
          </div>

          <div className="mt-12 grid gap-6 md:grid-cols-3">
            <Feature
              title="Face-based search"
              description="Find photos containing you without manually searching through hundreds or thousands of images."
            />

            <Feature
              title="Fast discovery"
              description="Search your collection using facial embeddings instead of relying only on filenames or folders."
            />

            <Feature
              title="Your collection"
              description="Keep your photos organized around the people who actually matter to you."
            />
          </div>
        </div>
      </section>

      {/* CTA */}
      <section className="border-t">
        <div className="mx-auto max-w-7xl px-6 py-24 text-center">
          <h2 className="text-3xl font-bold tracking-tight">
            Ready to find your photos?
          </h2>

          <p className="mx-auto mt-4 max-w-xl text-muted-foreground">
            Create your FaceFind account and start discovering photos of
            yourself.
          </p>

          <Button
            variant="default"
            className="mt-8 h-11 px-6 text-sm font-semibold"
            onClick={handleSignup}
          >
            Create account
          </Button>
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t">
        <div className="mx-auto flex max-w-7xl items-center justify-between px-6 py-6">
          <p className="text-sm text-muted-foreground">
            © {new Date().getFullYear()} FaceFind
          </p>

          <Button
            variant="link"
            className="text-sm font-medium"
            onClick={handleLogin}
          >
            Log in
          </Button>
        </div>
      </footer>
    </main>
  );
}

function Feature({
  title,
  description,
}: {
  title: string;
  description: string;
}) {
  return (
    <div className="rounded-xl border bg-background p-6">
      <div className="mb-4 flex h-10 w-10 items-center justify-center rounded-lg bg-primary/10">
        <div className="h-2.5 w-2.5 rounded-full bg-primary" />
      </div>

      <h3 className="font-semibold">{title}</h3>

      <p className="mt-2 text-sm leading-6 text-muted-foreground">
        {description}
      </p>
    </div>
  );
}
