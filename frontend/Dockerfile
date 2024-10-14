FROM node:18-alpine


WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 5173

# need to include --host because this exposes it to access outside the localhost container
CMD ["npm", "run", "dev", "--" ,"--host"] 
