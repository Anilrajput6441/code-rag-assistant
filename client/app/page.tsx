"use client";

import { useState, useEffect } from "react";
import Header from "@/components/Header";
import AuthBox from "@/components/AuthBox";
import IngestBox from "@/components/IngestBox";
import ChatBox from "@/components/ChatBox";

export default function Home() {
  const [activePanel, setActivePanel] = useState(0);

  const panels = [
    {
      title: "🔐 Authentication",
      component: <AuthBox />,
      info: (
        <div className="mt-6">
          <h3 className="font-semibold mb-3 text-white">How it works</h3>
          <ol className="list-decimal list-inside text-white/80 space-y-2 text-sm">
            <li>Login and add your Gemini API key</li>
            <li>Ingest a GitHub repository</li>
            <li>Ask questions and get grounded answers</li>
          </ol>
        </div>
      )
    },
    {
      title: "📁 Repository Ingestion",
      component: <IngestBox />
    },
    {
      title: "💬 Ask Questions",
      component: <ChatBox />
    }
  ];

  // Handle wheel scroll
  useEffect(() => {
    const handleWheel = (e: WheelEvent) => {
      e.preventDefault();
      if (e.deltaY > 0 && activePanel < panels.length - 1) {
        setActivePanel(prev => prev + 1);
      } else if (e.deltaY < 0 && activePanel > 0) {
        setActivePanel(prev => prev - 1);
      }
    };

    window.addEventListener('wheel', handleWheel, { passive: false });
    return () => window.removeEventListener('wheel', handleWheel);
  }, [activePanel, panels.length]);

  return (
    <div className="relative overflow-hidden">
      {/* Background Video */}
      <div className="fixed inset-0 z-0">
        <video 
          className="w-full h-full object-cover" 
          src="Rag-video.mp4" 
          autoPlay 
          muted
          playsInline
        ></video>
      </div>

      {/* Header */}
      <div className="relative z-10">
        <Header />
      </div>

      {/* Curved Carousel Container */}
      <div className="relative z-10 h-[80vh] flex items-center justify-center">
        <div 
          className="relative w-full h-full"
          style={{ perspective: "1200px" }}
        >
          {panels.map((panel, index) => {
            const offset = index - activePanel;
            const absOffset = Math.abs(offset);
            
            // Calculate positions for curved layout
            const translateX = offset * 400; // Horizontal spacing
            const translateZ = -absOffset * 200; // Depth
            const rotateY = offset * 25; // Rotation angle
            const scale = 1 - absOffset * 0.2; // Scale down distant panels
            const opacity = 1 - absOffset * 0.3; // Fade distant panels
            
            return (
              <div
                key={index}
                className={`absolute inset-0 flex items-center justify-center transition-all duration-700 ease-out ${
                  index === activePanel ? "z-20" : "z-10"
                }`}
                style={{
                  transform: `translateX(${translateX}px) translateZ(${translateZ}px) rotateY(${rotateY}deg) scale(${scale})`,
                  opacity: absOffset > 1 ? 0 : opacity,
                  transformStyle: "preserve-3d"
                }}
              >
                <div className="w-full max-w-2xl">
                  <div className="bg-white/10 backdrop-blur-md rounded-3xl p-8 border border-white/20 h-[60vh] overflow-y-auto shadow-2xl">
                    <h2 className="text-2xl font-semibold mb-6 text-white text-center">
                      {panel.title}
                    </h2>
                    <div className="space-y-4">
                      {panel.component}
                      {panel.info}
                    </div>
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      </div>

      {/* Navigation Dots */}
      <div className="relative z-10 flex justify-center gap-3 mt-8">
        {panels.map((_, index) => (
          <button
            key={index}
            onClick={() => setActivePanel(index)}
            className={`w-3 h-3 rounded-full transition-all duration-300 ${
              index === activePanel
                ? "bg-white scale-125"
                : "bg-white/40 hover:bg-white/60"
            }`}
          />
        ))}
      </div>

      {/* Navigation Arrows */}
      <div className="fixed left-8 top-1/2 transform -translate-y-1/2 z-20">
        <button
          onClick={() => setActivePanel(Math.max(0, activePanel - 1))}
          disabled={activePanel === 0}
          className="w-12 h-12 bg-white/10 backdrop-blur-md rounded-full flex items-center justify-center text-white text-xl disabled:opacity-30 disabled:cursor-not-allowed hover:bg-white/20 transition-all"
        >
          ←
        </button>
      </div>
      
      <div className="fixed right-8 top-1/2 transform -translate-y-1/2 z-20">
        <button
          onClick={() => setActivePanel(Math.min(panels.length - 1, activePanel + 1))}
          disabled={activePanel === panels.length - 1}
          className="w-12 h-12 bg-white/10 backdrop-blur-md rounded-full flex items-center justify-center text-white text-xl disabled:opacity-30 disabled:cursor-not-allowed hover:bg-white/20 transition-all"
        >
          →
        </button>
      </div>

      {/* Scroll Hint */}
      <div className="fixed bottom-8 left-1/2 transform -translate-x-1/2 z-20 text-white/60 text-sm text-center">
        <div className="flex items-center gap-2">
          <span>Scroll to navigate</span>
          <div className="w-6 h-10 border-2 border-white/40 rounded-full flex justify-center">
            <div className="w-1 h-3 bg-white/60 rounded-full mt-2 animate-bounce"></div>
          </div>
        </div>
      </div>
    </div>
  );
}