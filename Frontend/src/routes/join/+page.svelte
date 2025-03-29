<script>
    import CompetiterCard from "../../components/CompetiterCard/CompetiterCard.svelte";
</script>

<div
    class="flex flex-col items-center justify-center"
    style="background-image: url('src/lib/assets/image/play-bg.png'); background-size: cover; background-position: center; height: 100vh;"
>
<CompetiterCard
    profileImage="src/lib/assets/image/profile.jpg"
    playerName="John Doe"
    money={1000}
    happiness={80}
    knowledge={90}
    relationship={70}
    age0_12={[
        { type: "money", description: "Learning and growing." },
        { type: "happiness", description: "Exploring and discovering." },
        { type: "relationship", description: "Building a career and relationships." },
        { type: "knowledge", description: "Learning and growing." },
        // { type: "knowledge", description: "Exploring and discovering." },
        // { type: "relationship", description: "Building a career and relationships." },
        // { type: "money", description: "Learning and growing." },
        // { type: "knowledge", description: "Exploring and discovering." },
        
    ]}
    age13_18={[
        { type: "money", description: "Learning and growing." },
        { type: "knowledge", description: "Exploring and discovering." },
        { type: "relationship", description: "Building a career and relationships." }
    ]}
/>
</div>