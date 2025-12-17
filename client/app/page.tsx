import IngestBox from "@/components/IngestBox";
import ChatBox from "@/components/ChatBox";
import AuthBox from "@/components/AuthBox";

export default function Home() {
  return (
    <main className="max-w-3xl mx-auto p-6">
      <h1 className="text-2xl font-bold mb-6">
        AI Code Intelligence Assistant
      </h1>
      <AuthBox />
      <IngestBox />
      <ChatBox />
    </main>
  );
}
