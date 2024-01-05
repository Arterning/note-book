
in middleware.ts, we define middleware to protext al route except upload thing api
```jsx
import { authMiddleware } from "@clerk/nextjs";
 
// This example protects all routes including api/trpc routes
// Please edit this to allow other routes to be public as needed.
// See https://clerk.com/docs/references/nextjs/auth-middleware for more information about configuring your middleware
export default authMiddleware({
  publicRoutes: ["/api/uploadthing"]
});
 
export const config = {
  matcher: ["/((?!.*\\..*|_next).*)", "/", "/(api|trpc)(.*)"],
};

```

That ensure our application page only available when user is logged in 

When user logged in, he should first create server with name and cover iamges

![[Pasted image 20231231154104.png]]

The create server model is defined in `ModalProvider`, this comonent defined all the modal our application that needed.

![[Pasted image 20231231155528.png]]


The SetupPage is the first page, it is a server component, the logic is simple
find the fist server and if it exist , just go into the server, otherwise go to the InitialModal, 
InitialModal is the component that create server for user.

```jsx
import { redirect } from "next/navigation";

import { db } from "@/lib/db";
import { initialProfile } from "@/lib/initial-profile";
import { InitialModal } from "@/components/modals/initial-modal";

const SetupPage = async () => {
  const profile = await initialProfile();

  const server = await db.server.findFirst({
    where: {
      members: {
        some: {
          profileId: profile.id
        }
      }
    }
  });

  if (server) {
    return redirect(`/servers/${server.id}`);
  }

  return <InitialModal />;
}
 
export default SetupPage;
```

after server created, we ill goto the main page
![[Pasted image 20231231154903.png]]

## Server page

layout is three column

```jsx
import { NavigationSidebar } from "@/components/navigation/navigation-sidebar";

const MainLayout = async ({
  children
}: {
  children: React.ReactNode;
}) => {
  return ( 
    <div className="h-full">
      <div className="hidden md:flex h-full w-[72px] z-30 flex-col fixed inset-y-0">
        <NavigationSidebar />
      </div>
      <main className="md:pl-[72px] h-full">
        {children}
      </main>
    </div>
   );
}
 
export default MainLayout;
```

![[Pasted image 20231231163203.png]]

Then is the server Layout

```jsx
import { redirectToSignIn } from "@clerk/nextjs";
import { redirect } from "next/navigation";

import { db } from "@/lib/db";
import { currentProfile } from "@/lib/current-profile";
import { ServerSidebar } from "@/components/server/server-sidebar";

const ServerIdLayout = async ({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { serverId: string };
}) => {
  const profile = await currentProfile();

  if (!profile) {
    return redirectToSignIn();
  }

  const server = await db.server.findUnique({
    where: {
      id: params.serverId,
      members: {
        some: {
          profileId: profile.id
        }
      }
    }
  });

  if (!server) {
    return redirect("/");
  }

  return ( 
    <div className="h-full">
      <div 
      className="hidden md:flex h-full w-60 z-20 flex-col fixed inset-y-0">
        <ServerSidebar serverId={params.serverId} />
      </div>
      <main className="h-full md:pl-60">
        {children}
      </main>
    </div>
   );
}
 
export default ServerIdLayout;
```