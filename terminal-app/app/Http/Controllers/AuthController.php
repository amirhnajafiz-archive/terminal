<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Redis;
use TheSeer\Tokenizer\Token;

class AuthController extends Controller
{
    public function login(Request $request): JsonResponse
    {
        $name = $request->get('name');

        if (User::query()->where('name', '=', $name)->exists()) {
            $token = (new Token(1, $name, $name));
            Redis::set($name, $token);
        } else {
            $token = 'None';
        }

        return response()
            ->json([
                'token' => $token
            ]);
    }

    public function logout(Request $request): JsonResponse
    {
        $name = $request->get('name');
        $token = $request->get('token');

        if (Redis::exists($name) && Redis::get($name) == $token)
        {
            Redis::del($name);
        }

        return response()
            ->json([
                'status' => 'logged out'
            ]);
    }
}
