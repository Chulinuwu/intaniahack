import json
from diffusers import StableDiffusionPipeline
import torch

# Load JSON data from file
json_file = "events-data.json"  # Replace with your actual JSON file path
with open(json_file, "r", encoding="utf-8") as file:
    data = json.load(file)

# Initialize the Stable Diffusion pipeline
model_id = "sd-legacy/stable-diffusion-v1-5"
pipe = StableDiffusionPipeline.from_pretrained(model_id, torch_dtype=torch.float16)
pipe = pipe.to("cuda")

# Loop through each item in the JSON and generate an image
for item in data:
    if item['id'] != "age1_8":
        continue
    # 2 7
    # 1 8
    prompt = item["description"]  # Use description as the prompt
    prompt = "You are " + prompt +" Gen as cartoon art for game safe for kids"
    image = pipe(prompt).images[0]  # Generate the image
    
    # Save the image using the 'id' as the filename
    filename = f"{item['id']}.png"
    image.save(filename)
    print(f"Saved: {filename}")